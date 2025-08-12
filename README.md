<p align="center"><img src="https://github.com/JayyDoesDev/kiku/blob/main/.github/assets/kiku.png"></p>
<h1 align="center">Kiku</h1>
<h2 align="center">💐🌻Kiku, a simple sharex server for uploading images and shortening urls!</h2>
<div>
  <h2 align="center">
    <img src="https://img.shields.io/github/commit-activity/m/jayydoesdev/kiku">
    <img src="https://img.shields.io/github/license/jayydoesdev/kiku">
    <img src="https://img.shields.io/github/languages/top/jayydoesdev/kiku">
    <img src="https://img.shields.io/github/contributors/jayydoesdev/kiku">
    <img src="https://img.shields.io/github/last-commit/jayydoesdev/kiku">
  </h2>
</div>

- 🖼️ **Image Sharing** - Stores images and allows you to share them
- 🔗 **Url Shortening** - Creates shortened urls for any url you desire
- 🔑 **Authorization** API Key authorization 

## Project Structure
```
├── kiku/
│   ├── db/         #  Database functions
│   ├── lib/        # Utilities and database connection
│   ├── middleware/ # Middleware in which contains auth
│   ├── public/     # HTML and CSS files
│   ├── routes/     # Backend routes
│   └── storage/    # Image storage
└── main.go
```