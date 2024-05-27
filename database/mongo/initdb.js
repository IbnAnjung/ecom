db.createUser(
  {
      user: "edot",
      pwd: "edot_secret",
      roles: [
          {
              role: "readWrite",
              db: "edot"
          }
      ]
  }
);

db = new Mongo().getDB("edot");

db.createCollection("product")

db.product.createIndex({"name": "text", "description": "text"})

db.product.insertMany([
  {
    _id: ObjectId("6651a98294344ee37278f6d7"),
    name: "New Orleans Coffee",
    description: "Cafe Noir from New Orleans is a spiced, nutty coffee made with chicory.",
    price: 50000,
    store_id: 1,
  },
  {
    _id: ObjectId("6651a98294344ee37278f6d8"),
    name: "Affogato",
    description: "An Italian sweet dessert coffee made with fresh-brewed espresso and vanilla ice cream.",
    price: 25000,
    store_id: 1,
  },
  {
    _id: ObjectId("6651a98294344ee37278f6d6"),
    name: "Cafecito",
    description: "A sweet and rich Cuban hot coffee made by topping an espresso shot with a thick sugar cream foam.",
    price: 500000,
    store_id: 1,
  },
  {
    _id: ObjectId("6651a98294344ee37278f6d9"),
    name: "Maple Latte",
    description: "A wintertime classic made with espresso and steamed milk and sweetened with some maple syrup.",
    price: 50200,
    store_id: 2,
  },
  {
    _id: ObjectId("6651a98294344ee37278f6da"),
    name: "Pumpkin Spice Latte",
    description: "It wouldn't be autumn without pumpkin spice lattes made with espresso, steamed milk, cinnamon spices, and pumpkin puree.",
    price: 10000,
    store_id: 2,
  }
]);
