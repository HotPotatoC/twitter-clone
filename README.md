<br />
<p align="center">
  <a href="https://github.com/HotPotatoC/heiver">
    <img src=".github/assets/Twitter-clone.png" alt="Logo">
  </a>

  <h3 align="center">Twitter Clone</h3>

  <p align="center">A Twitter clone created with Golang, PostgreSQL, Redis, VueJS and Vite with support for dark mode and light mode using TailwindCSS
  </p>

  <p align="center">Status: Some features are yet to be implemented</p>
</p>

---

# Preview

![preview](.github/assets/twitterclone.gif)

*For more check out some screenshots [here](SCREENSHOTS.md)*

# Features ✨

> NOTE: Not all features from twitter are implemented because of how big Twitter is, Only the main features are implemented atm

- Modular Architecture
- Database migration tool using [migrate](https://github.com/golang-migrate/migrate)
- Golang Hot-reloading using [air](https://github.com/cosmtrek/air)
- Supports dark-mode and light-mode with [TailwindCSS](http://tailwindcs.com/)
- Database seeding script using NodeJS
- Authentication using JWT Refresh token flow and Redis for token blacklisting
- Strongly typed Vuex store
- List Tweets feed
- Create Tweets with images
- Retweets
- Reply to Tweets or reply to another reply!
- Like Tweets
- Follow users
- Images & Media uploads stored in AWS S3 Buckets
- Up to 4 images in a single tweet with the same layout as Twitter
- Crop profile image
- Edit Profile Details
- Edit Profile Image
- See who a user is following and see their followers

# Tech 🛠

- [Golang](golang.org)
- [Fiber HTTP framework](https://github.com/gofiber/fiber)
- [PostgreSQL](postgresql.org)
- [Redis](redis.io)
- [NodeJS](https://nodejs.org/en/)
- [TypeScript](https://www.typescriptlang.org/)
- [migrate](https://github.com/golang-migrate/migrate)
- [air](https://github.com/cosmtrek/air)
- [Amazon Web Service S3](https://aws.amazon.com/s3/)
- [Vue 3](https://v3.vuejs.org/)
- [Vite 2.0](https://vitejs.dev/)
- [Vuex 4](https://next.vuex.vuejs.org)
- [Vue Router 4](https://next.router.vuejs.org)
- [TailwindCSS](http://tailwindcs.com/)

# How to run locally

Check [here](RUNNING_LOCALLY.md) on how to run locally

# Resources & references used

- https://twitter.com
- https://about.twitter.com/en/who-we-are/brand-toolkit
- https://github.com/shuber/postgres-twitter
- [Build a twitter clone using vue.js and tailwind css! (by: this.stephie)](https://www.youtube.com/watch?v=bQU-jPyQJ4A)
- [Is SELECT * Expensive? (by: Hussein Nasser)](https://www.youtube.com/watch?v=QQVNVOneZNg)
- [SELECT COUNT (*) can impact your Backend Application Performance, here is why (by: Hussein Nasser)](https://www.youtube.com/watch?v=8xKS7QQKgzk)
- [Full Text Search PostgreSQL (by: Ben Awad)](https://www.youtube.com/watch?v=szfUbzsKvtE)
- https://www.postgresql.org/message-id/20050810133157.GA46247@winnie.fuhr.org
- https://dev.to/shubhadip/vue-3-vuex-4-modules-typescript-2i2o
- https://dev.to/3vilarthas/vuex-typescript-m4j
- [Today i learned golang live reload for development using docker compose air (by: Iman Tumorang)](https://medium.com/easyread/today-i-learned-golang-live-reload-for-development-using-docker-compose-air-ecc688ee076)

# Improvements

- [ ] Hashtags and mentions
- [ ] Notifications
- [ ] Bookmarks
- [ ] Lists
- [ ] Trending section

# Disclaimer

twitter-clone is created for educational purposes only.

I do not work for Twitter nor Twitter the company itself has any associations / involvements in this project.