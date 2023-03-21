db = new Mongo().getDB("carpool");
db.users.drop();
db.users.createIndex( { "email": 1 }, { unique: true } )
db.users.createIndex( { "phone": 1 }, { unique: true } )
db.users.insertMany([
    {
        "_id": ObjectId("5fb7b1d521e4f91673dc2293"),
        "name": "First User",
        "email": "f@f.f",
        "phone": "1234567890",
        "password": "��I\u0000+Dܢ�r!:Z�\u000e0\""
    }
])