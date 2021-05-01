![Twitter clone](.github/Twitter-clone.png)

# Twitter clone

> A Twitter clone created with Golang, PostgreSQL, Redis, VueJS and Vite with support for dark mode and light mode using TailwindCSS

# Features ✨

Only the main features are implemented atm

- Modular Architecture
- Database migration tool using [migrate](https://github.com/golang-migrate/migrate)
- Database seeding script using NodeJS
- Authentication using JWT Refresh token flow and Redis for JWT blacklisting
- Strongly typed Vuex store
- List Tweets feed
- Create Tweets
- Reply to Tweets or reply to another reply!
- Like Tweets
- Follow users
- Unfollow users
- Images & Media uploading using AWS S3 Buckets
- Edit Profile Details
- Edit Profile Image
- See who a user is following and see their followers

# Tech 🛠

- [Golang](golang.org)
- [Fiber HTTP framework](https://github.com/gofiber/fiber)
- [PostgreSQL](postgresql.org)
- [Redis](redis.io)
- [TypeScript](https://www.typescriptlang.org/)
- [Vue 3](https://v3.vuejs.org/)
- [Vite 2.0](https://vitejs.dev/)
- [Vuex 4](https://next.vuex.vuejs.org)
- [Vue Router 4](https://next.router.vuejs.org)
- [TailwindCSS](http://tailwindcs.com/)

# Installation - Running locally 💻

> NOTE: To run this app locally you need to have an AWS S3 Bucket available so that uploading images will work. At the time im writing this I'm using an account provided by AWS educate which is an available option for students or you can start a free trial if you have a credit card.

1. **Clone the repository**

```sh
❯ git clone https://github.com/HotPotatoC/twitter-clone.git

❯ cd twitter-clone
```

2. **Create .env file in `configs/` directory by copying `configs/.env.example` and setup the environment variables**

3. **Create the PostgreSQL database and run migrations**

> To run the migrations first install the **migrate** tool [here](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

```sh
❯ createdb twitterclone

❯ ./scripts/run_migrations.sh Your_PostgreSQL_URL
```

1. **Run the backend server**

```sh
❯ go run cmd/rest/main.go --prefork
```

5. **On another terminal tab setup the frontend**

```sh
❯ cd web
```

Using yarn

```sh
❯ yarn
```

Using npm

```sh
❯ npm install
```

6. **Run the frontend**

```sh
❯ yarn dev
```

And there you go!

# Todo

SS = Server Side
CS = Client Side

- [x] Create reply (CS)
- [x] Logout (CS)
- [x] Favorite tweet (CS)
- [x] Profile View (CS)
- [x] Only see followed user Tweets in the feed (SS & CS)
- [ ] Avatar image (SS & CS)
  - [x] Update profile image (SS)
  - [ ] Update profile image (CS)
  - [ ] Crop image (CS)
  - [ ] Lazy load (CS)
- [ ] Attach an image to a Tweet (SS & CS)
- [x] 'Replying to ...' design Tweet card (CS)
- [ ] List profile followers and followings (CS)
- [x] Support for link parsing on Tweet's content (CS)
- [ ] Retweets (SS & CS)
- [ ] Profile Tweets & replies tab (CS)
- [ ] List user likes / Profile likes tab (SS & CS)
- [ ] Profile media tab (CS)
- [ ] Pagination on search results (SS)
- [ ] Toggle dark mode and light mode (CS)

# Improvements

- [ ] Tweet threads
- [ ] Hashtags and mentions
- [ ] Notifications with Redis pubsub & websockets (?)
- [ ] Bookmarks
- [ ] Lists
- [ ] Trending section

# Disclaimer

`twitter-clone` is created for educational purposes only. I do not work for Twitter nor Twitter the copyright holder have any associations with this experiment.