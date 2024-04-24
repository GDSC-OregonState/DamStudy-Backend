let username = "damstudy";
let password = "damstuy";

if (typeof username === "undefined" || username === null) {
  throw new Error("The username is missing.");
}

if (typeof password === "undefined" || password === null) {
  throw new Error("The password is missing.");
}

db.createUser({
  user: username,
  pwd: password,
  roles: [{ role: "readWrite", db: "damstudy" }],
});

db.createCollection("rooms");

let damstudyCollection =
  db.getCollection("rooms") || db.createCollection("rooms");

damstudyCollection.insertOne({
  name: "KEC 1001 Study Room",
  image:
    "https://egis.umn.edu/studyspace_v2/studyspaceimages/10ChurchStreet-101.jpg",
  noiseLevel: "Quiet",
  seats: 4,
  technology: ["Whiteboard", "Projector"],
  seating: "Table",
  location: "Kelley Engineering Center",
  coordinates: {
    latitude: 44.56706903872953,
    longitude: -123.27873483900234,
  },
});
