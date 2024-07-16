const express = require('express');
const { Pool } = require('pg');
const bodyParser = require('body-parser');

// PostgreSQL connection configuration
const pool = new Pool({
  host: 'localhost',
  port: 5432,
  user: 'kiran',
  password: 'kiran0612',
  database: 'testdb',
});

// Struct to model the PostgreSQL table
const app = express();
app.use(bodyParser.json());

app.listen(8000, () => {
  console.log('Server is running on port 8000');
});

// Create a new user
app.post('/users', async (req, res) => {
  const { name, age, gender, email, mobile, address } = req.body;
  try {
    const result = await pool.query(
      'INSERT INTO users (name, age, gender, email, mobile, address) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id',
      [name, age, gender, email, mobile, address]
    );
    res.json({ id: result.rows[0].id, ...req.body });
  } catch (err) {
    console.error('Failed to create user:', err);
    res.status(500).json({ error: 'Failed to create user' });
  }
});

// Bulk create users
app.post('/users/bulk', async (req, res) => {
  const users = req.body;
  const client = await pool.connect();
  try {
    await client.query('BEGIN');
    const queryText = 'INSERT INTO users(name, age, gender, email, mobile, address) VALUES($1, $2, $3, $4, $5, $6)';
    for (const user of users) {
      await client.query(queryText, [user.name, user.age, user.gender, user.email, user.mobile, user.address]);
    }
    await client.query('COMMIT');
    res.json(users);
  } catch (err) {
    await client.query('ROLLBACK');
    console.error('Failed to execute bulk insert:', err);
    res.status(500).json({ error: 'Failed to execute bulk insert' });
  } finally {
    client.release();
  }
});

// Get a user by ID
app.get('/users/:id', async (req, res) => {
  const id = parseInt(req.params.id);
  try {
    const result = await pool.query('SELECT id, name, age, gender, email, mobile, address FROM users WHERE id=$1', [id]);
    if (result.rows.length === 0) {
      res.status(404).json({ error: 'User not found' });
    } else {
      res.json(result.rows[0]);
    }
  } catch (err) {
    console.error('Failed to get user:', err);
    res.status(500).json({ error: 'Failed to get user' });
  }
});

// Update a user by ID
app.put('/users/:id', async (req, res) => {
  const id = parseInt(req.params.id);
  const { name, age, gender, email, mobile, address } = req.body;
  try {
    await pool.query(
      'UPDATE users SET name=$2, age=$3, gender=$4, email=$5, mobile=$6, address=$7 WHERE id=$1',
      [id, name, age, gender, email, mobile, address]
    );
    res.json({ id, ...req.body });
  } catch (err) {
    console.error('Failed to update user:', err);
    res.status(500).json({ error: 'Failed to update user' });
  }
});

// Delete a user by ID
app.delete('/users/:id', async (req, res) => {
  const id = parseInt(req.params.id);
  try {
    await pool.query('DELETE FROM users WHERE id=$1', [id]);
    res.json({ message: 'User deleted' });
  } catch (err) {
    console.error('Failed to delete user:', err);
    res.status(500).json({ error: 'Failed to delete user' });
  }
});
