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
- 🔑 **Authorization** - API Key authorization

---

## 🖼️ Preview
### Web Page
<p align="center"><img src="https://github.com/jayydoesdev/kiku/blob/main/.github/assets/preview_1.png?raw=true" alt="Web Page Preview" width="700"/></p>

### Image Sharing
<p align="center"><img src="https://github.com/jayydoesdev/kiku/blob/main/.github/assets/preview_2.png?raw=true" alt="Image Sharing Preview" width="700"/></p>

### Url Shortening
<p align="center"><img src="https://github.com/jayydoesdev/kiku/blob/main/.github/assets/preview_3.png?raw=true" alt="Url Shortening Preview" width="700"/></p>

---

## 📁 Project Structure
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

---

## 🚀 Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/jayydoesdev/kiku.git
cd kiku
```

### 2. Set up your environment

Create a `.env` file in the root:

```env
API_URL=your_api_url
PORT=:port
TOKEN=your_api_token
DATABASE_URL=mongodb_url
```
### 3. Run the bot

```bash
go run main.go
```

---

## 📄 License

This project is licensed under the MIT License — see [LICENSE](LICENSE) for details.

<a href="https://github.com/jayydoesdev/kiku/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=jayydoesdev/kiku" />
</a>

