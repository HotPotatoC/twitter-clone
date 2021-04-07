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

  for (let i = 0; i < n; i++) {
    const randomUser = await pool.query(
      "select id from users order by random() limit 1"
    );
    const randomTweet = await pool.query(
      "select id from tweets order by random() limit 1"
    );
    const insertedReplyTweet = await pool.query(
      "INSERT INTO tweets(content, id_user, created_at) VALUES($1, $2, $3) RETURNING id",
      [faker.hacker.phrase(), Number(randomUser.rows[0].id), faker.date.past()]
    );
    await pool.query("INSERT INTO replies(id_reply, id_tweet) VALUES($1, $2)", [
      Number(insertedReplyTweet.rows[0].id),
      Number(randomTweet.rows[0].id),
    ]);
  }
  await pool.end();
  console.log(
    `Successfully seeded the replies table with ${n} records`
  );
}

main();
