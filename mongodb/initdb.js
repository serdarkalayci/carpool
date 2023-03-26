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