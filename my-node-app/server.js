const express = require('express');
const { MongoClient, ObjectId } = require('mongodb');
const bodyParser = require('body-parser');

// MongoDB connection URI
const connectionString = 'mongodb+srv://kiransai:Kiran0612@cluster0.4nztuz7.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0';

// Database and collection names
const dbName = 'testdb';
const colName = 'users';

const app = express();
app.use(bodyParser.json());

let collection;

// Initialize MongoDB client
MongoClient.connect(connectionString, { useNewUrlParser: true, useUnifiedTopology: true })
  .then(client => {
    console.log('Connected to MongoDB!');
    const db = client.db(dbName);
    collection = db.collection(colName);

    // Start server
    app.listen(8000, () => {
      console.log('Server is running on port 8000');
    });
  })
  .catch(err => {
    console.error('Failed to connect to MongoDB:', err);
    process.exit(1);
  });

// Create a new user
app.post('/users', async (req, res) => {
  const user = req.body;
  user._id = new ObjectId(); // Generate a new ID

  try {
    await collection.insertOne(user);
    res.json(user);
  } catch (err) {
    console.error('Failed to create user:', err);
    res.status(500).json({ error: 'Failed to create user' });
  }
});

// Bulk create users
app.post('/users/bulk', async (req, res) => {
  const users = req.body.map(user => ({ ...user, _id: new ObjectId() })); // Generate new IDs for each user

  try {
    await collection.insertMany(users);
    res.json(users);
  } catch (err) {
    console.error('Failed to create users:', err);
    res.status(500).json({ error: 'Failed to create users' });
  }
});

// Get a user by ID
app.get('/users/:id', async (req, res) => {
  const id = req.params.id;

  try {
    const user = await collection.findOne({ _id: new ObjectId(id) });
    if (!user) {
      res.status(404).json({ error: 'User not found' });
    } else {
      res.json(user);
    }
  } catch (err) {
    console.error('Failed to get user:', err);
    res.status(500).json({ error: 'Failed to get user' });
  }
});

// Update a user by ID
app.put('/users/:id', async (req, res) => {
  const id = req.params.id;
  const update = { $set: req.body };

  try {
    const result = await collection.updateOne({ _id: new ObjectId(id) }, update);
    if (result.matchedCount === 0) {
      res.status(404).json({ error: 'User not found' });
    } else {
      res.json(req.body);
    }
  } catch (err) {
    console.error('Failed to update user:', err);
    res.status(500).json({ error: 'Failed to update user' });
  }
});

// Delete a user by ID
app.delete('/users/:id', async (req, res) => {
  const id = req.params.id;

  try {
    const result = await collection.deleteOne({ _id: new ObjectId(id) });
    if (result.deletedCount === 0) {
      res.status(404).json({ error: 'User not found' });
    } else {
      res.json({ message: 'User deleted' });
    }
  } catch (err) {
    console.error('Failed to delete user:', err);
    res.status(500).json({ error: 'Failed to delete user' });
  }
});
