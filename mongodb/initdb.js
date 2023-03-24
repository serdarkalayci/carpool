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
db.geography.drop();
db.geography.insertMany([
    {
        _id:   ObjectId("666f6f2d6261722d71757578"),
        name: "United States",
        cities: [
            {
                name:   "New York City",
                ballot: true
            },
            {
                name:   "Los Angeles",
                ballot: false
            },
            {
                name:   "Chicago",
                ballot: true
            }
        ]
    },
    {
        _id: ObjectId("0123456789ab0123456789ab"),
        name: "The Netherlands",
        cities: [
            {
                name:   "Amsterdam",
                ballot: true
            },
            {
                name: "Haarlem",
                ballot: false
            },
            {
                name: "Rotterdam",
                ballot: false
            },
            {   
                name: "Utrecht",
                ballot: false
            }
        ]
    }
])