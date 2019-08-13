using System;

using MongoDB.Bson;
using MongoDB.Bson.Serialization.Attributes;

namespace mongo.Models
{
    public class User
    {
        [BsonId]
        [BsonRepresentation(BsonType.ObjectId)]
        public string Id { get; set; }

        public string Firstname { get; set; }

        public string Lastname { get; set; }

        public int Age { get; set; }
    }
}
