const path = require("path");
const dotenv = require("dotenv");
const { Pool } = require("pg");
const faker = require("faker");

dotenv.config({
  path: path.normalize(path.join(__dirname, "../../configs/.env")),
});

async function main() {
  const n = Number(process.argv[2]);
  const pool = new Pool({
    user: process.env.DB_USER,
    password: process.env.DB_PASSWORD,
    host: process.env.DB_HOST,
    port: Number(process.env.DB_PORT),
    database: process.env.DB_DATABASE,
  });

  const query = `INSERT INTO users(name, email, password, created_at, bio, location, website, birth_date) VALUES($1, $2, $3, $4, $5, $6, $7, $8)`;
  for (let i = 0; i < n; i++) {
    await pool.query(query, [
      faker.internet.userName(),
      faker.internet.email(),
      "$2y$12$/HayGBMevM7EkbRho3xlC.nu/8.OeHGNi9iqGneNIDBrqnkD9fa1.",
      faker.date.past(),
      faker.hacker.phrase(),
      faker.address.country(),
      faker.internet.url(),
      faker.date.between("1990-01-01", "2000-01-01"),
    ]);
  }
  await pool.end();
  console.log(
    `Successfully seeded the users table with ${n} records (all users have the same password: 123123)`
  );
}

main();
