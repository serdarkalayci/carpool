db = new Mongo().getDB("carpool");
db.users.drop();
db.users.createIndex( { "email": 1 }, { unique: true } )
db.users.createIndex( { "phone": 1 }, { unique: true } )
db.users.insertMany([
    {
        "_id": ObjectId("5fb7b1d521e4f91673dc2293"),
        "name": "Serdar Kalaycı",
        "email": "s@s.s",
        "phone": "1234567890",
        "password": "\u0006\u0098á*GŸXwb)Z8¿7\u0087\u0091à*Ë"
    }
])

db.trips.drop();
db.trips.insertMany([
  {
    "_id": ObjectId("642064a2657e328de2a7d631"),
    "supplierid":ObjectId("5fb7b1d521e4f91673dc2293"),
    "countryid":ObjectId("641f01dfcc0d85792b29d254"),
    "origin":"The Hague",
    "destination":"Amsterdam",
    "stops":["Haarlem","Almere"],
    "tripdate": ISODate("2023-03-26T00:00:00Z"),
    "availableseats":3,
    "note":"not yok bu kez"
},
{
    "_id":ObjectId("642075b89e46e44d46172af2"),
    "supplierid":ObjectId("5fb7b1d521e4f91673dc2293"),
    "countryid":ObjectId("641f01dfcc0d85792b29d254"),
    "origin":"Utrecht",
    "destination":"Amsterdam",
    "stops":["Haarlem","Almere","Eindhoven"],
    "tripdate":ISODate("2023-03-26T00:00:00Z"),
    "availableseats":3,
    "note":"not yok bu kez"
},
{
    "_id":ObjectId("642075d59e46e44d46172af3"),
    "supplierid":ObjectId("5fb7b1d521e4f91673dc2293"),
    "countryid":ObjectId("641f01dfcc0d85792b29d254"),
    "origin":"Tilburg",
    "destination":"Amsterdam",
    "stops":["Haarlem","Almere","Eindhoven"],
    "tripdate":ISODate("2023-03-26T00:00:00Z"),
    "availableseats":3,
    "note":"not yok bu kez"
},
{
    "_id":ObjectId("642075f49e46e44d46172af4"),
    "supplierid":ObjectId("5fb7b1d521e4f91673dc2293"),
    "countryid":ObjectId("641f01dfcc0d85792b29d254"),
    "origin":"Tilburg",
    "destination":"Rotterdam",
    "stops":["Haarlem","Almere","Eindhoven"],
    "tripdate":ISODate("2023-03-26T00:00:00Z"),
    "availableseats":3,
    "note":"not yok bu kez"
}
])

db.createView( "tripdetail", "trips", [
  {
    $lookup:
      {
          from: "users",
          localField: "supplierid",
          foreignField: "_id",
          as: "userDocs"
      }
  },
  {
    $project:
      {
        _id: 1,
        countryid: 1,
        origin: 1,
        destination: 1,
        stops: 1,
        tripdate: 1,
        availableseats: 1,
        note: 1,
        username: "$userDocs.name"
      }
  },
  { 
    $unwind: "$username" 
  },
  {
    $lookup:
        {
          from: "geography",
          localField: "countryid",
          foreignField: "_id",
          as: "countryDocs"
        }
  },
  {
    $project:
        {
          _id: 1,
          origin: 1,
          destination: 1,
          stops: 1,
          tripdate: 1,
          availableseats: 1,
          username: 1,
          note: 1,
          countryname: "$countryDocs.name"
        }
  },
  { 
    $unwind: "$countryname" 
  }
] )

db.geography.drop();
db.geography.insertMany([
    {
        "_id": ObjectId("6420af53150f72f9e1feeb7c"),
        "name": "France",
        "cities": [
          {
            "name": "Paris",
            "ballot": false
          },
          {
            "name": "Marseille",
            "ballot": false
          },
          {
            "name": "Lyon",
            "ballot": false
          },
          {
            "name": "Toulouse",
            "ballot": false
          },
          {
            "name": "Nice",
            "ballot": false
          },
          {
            "name": "Nantes",
            "ballot": false
          },
          {
            "name": "Strasbourg",
            "ballot": false
          },
          {
            "name": "Montpellier",
            "ballot": false
          },
          {
            "name": "Bordeaux",
            "ballot": false
          },
          {
            "name": "Lille",
            "ballot": false
          },
          {
            "name": "Rennes",
            "ballot": false
          },
          {
            "name": "Reims",
            "ballot": false
          },
          {
            "name": "Le Havre",
            "ballot": false
          },
          {
            "name": "Saint-Étienne",
            "ballot": false
          },
          {
            "name": "Toulon",
            "ballot": false
          },
          {
            "name": "Grenoble",
            "ballot": false
          },
          {
            "name": "Dijon",
            "ballot": false
          },
          {
            "name": "Nîmes",
            "ballot": false
          },
          {
            "name": "Angers",
            "ballot": false
          },
          {
            "name": "Villeurbanne",
            "ballot": false
          },
          {
            "name": "Le Mans",
            "ballot": false
          },
          {
            "name": "Aix-en-Provence",
            "ballot": false
          },
          {
            "name": "Brest",
            "ballot": false
          },
          {
            "name": "Limoges",
            "ballot": false
          },
          {
            "name": "Tours",
            "ballot": false
          },
          {
            "name": "Amiens",
            "ballot": false
          },
          {
            "name": "Perpignan",
            "ballot": false
          },
          {
            "name": "Metz",
            "ballot": false
          },
          {
            "name": "Besançon",
            "ballot": false
          },
          {
            "name": "Boulogne-Billancourt",
            "ballot": false
          },
          {
            "name": "Orléans",
            "ballot": false
          }
        ]
      },
    {
        "_id":ObjectId("641f01dfcc0d85792b29d254"),
        "name": "The Netherlands",
        "cities": [
          {
            "name": "Amsterdam",
            "ballot": true
          },
          {
            "name": "Haarlem",
            "ballot": false
          },
          {
            "name": "Rotterdam",
            "ballot": true
          },
          {
            "name": "Utrecht",
            "ballot": false
          },
          {
            "name": "Groningen",
            "ballot": false
          },
          {
            "name": "Eindhoven",
            "ballot": false
          },
          {
            "name": "The Hague",
            "ballot": false
          },
          {
            "name": "Tilburg",
            "ballot": false
          },
          {
            "name": "Almere",
            "ballot": false
          },
          {
            "name": "Breda",
            "ballot": false
          },
          {
            "name": "Nijmegen",
            "ballot": false
          },
          {
            "name": "Enschede",
            "ballot": false
          },
          {
            "name": "Arnhem",
            "ballot": false
          },
          {
            "name": "Maastricht",
            "ballot": false
          },
          {
            "name": "Leeuwarden",
            "ballot": false
          },
          {
            "name": "Zwolle",
            "ballot": false
          },
          {
            "name": "Delft",
            "ballot": false
          },
          {
            "name": "Leiden",
            "ballot": false
          },
          {
            "name": "Amersfoort",
            "ballot": false
          },
          {
            "name": "Hilversum",
            "ballot": false
          },
          {
            "name": "Assen",
            "ballot": false
          },
          {
            "name": "Emmen",
            "ballot": false
          },
          {
            "name": "Gouda",
            "ballot": false
          },
          {
            "name": "Hoorn",
            "ballot": false
          },
          {
            "name": "Lelystad",
            "ballot": false
          },
          {
            "name": "Purmerend",
            "ballot": false
          },
          {
            "name": "Roosendaal",
            "ballot": false
          },
          {
            "name": "Schiedam",
            "ballot": false
          },
          {
            "name": "Spijkenisse",
            "ballot": false
          },
          {
            "name": "Veenendaal",
            "ballot": false
          },
          {
            "name": "Vlaardingen",
            "ballot": false
          },
          {
            "name": "Zaandam",
            "ballot": false
          },
          {
            "name": "Zeist",
            "ballot": false
          }
        ]
      }      
])