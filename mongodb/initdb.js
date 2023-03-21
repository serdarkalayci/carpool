db = new Mongo().getDB("carpool");
db.users.drop();
db.users.insertMany([
    {
        "_id": ObjectId("5fb7b1d521e4f91673dc2293"),
        "name": "First User",
        "username": "firstu",
        "password": "��I\u0000+Dܢ�r!:Z�\u000e0\""
    }
])