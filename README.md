[//]: # "header"
<h1 align="center">🏋️‍♂️Swole is the Goal🏋️‍♂️</h1>

[//]: # "tech stack used"
<div align="center">
   <img src="https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white" alt="Postgres" />
   <img src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white" alt="Go" />
   <img src="https://img.shields.io/badge/typescript-%23007ACC.svg?style=for-the-badge&logo=typescript&logoColor=white" alt="TypeScript" />
   <img src="https://img.shields.io/badge/github%20actions-%232671E5.svg?style=for-the-badge&logo=githubactions&logoColor=white" alt="GitHub Actions" />
   <img src="https://img.shields.io/badge/svelte-%23f1413d.svg?style=for-the-badge&logo=svelte&logoColor=white" alt="Svelte" />
</div>


[//]: # "catch"
<p align="center">
   A web-based, open-source alternative to track your lifts 🏋️
</p>

## Why was this made? 🤔💭
* Created as a solution to opting out of continuing to buy [Strong's subcription](https://www.strong.app/) to track my weightlifting journey
* During my time at University learning relational databases, I realized that the underlying technology to store and track progression within the gym was something that I would love to challenge myself with!
* An open-source solution to save money and customize to a communities liking 🤩💫

## Screenshots of the application! 😲🚀

## How do I run this locally? 💚🙂
> [!NOTE]\
> This program is fully hosted right now via localhost and package by a Docker container. I would love to host this on AWS, but I do not have the funds to manage their RDS instance :()
>
> Click -> [here](https://docs.docker.com/desktop/install/windows-install/) to go to the installation page.

With Docker Desktop installed and opened, you can run these commands at the root directory of this GitHub repo...
```shell
# if you are running this for the first time, or have made changes and want to see it take affect on its deployment, use this
docker compose up -d --build
# starts the application in its detached state
docker compose up -d
# turns off the application
docker compose down
# wipes the container clean, effectively deleting the volumes for the database
docker compose down -v
```
And then open up [localhost:4173](localhost:4173) 💚

## The Wiki 📕
Visit the [wiki](https://github.com/emoral435/swole-goal/wiki) in order to see more documentation about...
* How to use the software
* The database, and its general initial schema
* Contribution Guidelines
* Go's API documentation
* Dev dependencies for local development
* Local / Cloud usage
And more!

## Get in touch 💬
If you liked what you saw, feel free to contact me! email: emoral435@gmail.com

[Star Logs 🚀](https://starlogs.dev/emoral435/swole-goal)
