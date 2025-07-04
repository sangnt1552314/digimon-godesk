# 🦕 Digimon GoDesk

A Digimon information webapp built with Go to help me study Golang while exploring the digital world of [Digimon](https://en.wikipedia.org/wiki/Digimon)! 

## 🎯 Purpose

This project serves as my hands-on learning experience with Golang. What better way to learn a new language than building something about Digimon? After all, both involve evolution and transformation! 🚀

## 😅 A Funny Confession

**Full transparency:** I don't know CSS, and GitHub Copilot has been my styling savior throughout this project. Every beautiful button, every aligned div, every responsive layout - all thanks to AI assistance. If you see any decent styling, know that it wasn't me! 🎨✨

## 🌐 Data Source

This project utilizes the fantastic [Digi-API](https://digi-api.com/) to fetch Digimon information. Huge thanks to the maintainers for providing such a comprehensive and well-structured API for Digimon data!

## 🏗️ Project Structure

```
digimon-godesk/
├── cmd/                    # Application entry point
│   └── main.go
├── internal/              # Private application code
│   ├── config/           # Configuration management
│   ├── handler/          # HTTP handlers
│   │   ├── api/         # API endpoints
│   │   ├── middleware/  # HTTP middleware
│   │   └── web/         # Web page handlers
│   ├── logger/          # Logging utilities
│   ├── models/          # Data models
│   ├── services/        # Business logic
│   └── utils/           # Utility functions
├── templates/            # HTML templates
├── assets/              # Static assets (CSS, JS, images)
├── storage/             # Application storage
└── docker-compose.yml   # Docker setup
```

## 🚀 Getting Started

### Prerequisites

- Go 1.23 or higher
- Docker & Docker Compose (optional)

### Running Locally

1. Clone the repository:
```bash
git clone https://github.com/sangnt1552314/digimon-godesk.git
cd digimon-godesk
```

2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run cmd/main.go
```

4. Open your browser and navigate to `http://localhost:8080`

### Running with Docker

```bash
docker-compose up --build
```

## 🔧 Features

- 📱 Responsive web interface (thanks to AI styling!)
- 🔍 Search and browse Digimon
- 📊 Display detailed Digimon information
- 🌐 RESTful API endpoints
- 📝 Comprehensive logging
- 🐳 Docker support

## 🧠 What I'm Learning

Through this project, I'm exploring:

- Go web frameworks and HTTP handling
- Project structure and organization in Go
- API integration and HTTP clients
- Template rendering and web development
- Docker containerization
- Middleware implementation
- Logging and error handling
- And most importantly... how to make things look good without knowing CSS! 😂

## 🦖 About Digimon

[Digimon](https://en.wikipedia.org/wiki/Digimon) (Digital Monsters) is a Japanese media franchise encompassing virtual pet toys, anime, manga, video games, films and a trading card game. The franchise focuses on creatures called "Digimon" who inhabit a "Digital World," parallel to the real world.

Just like how Digimon evolve and grow stronger, I hope my Go skills evolve through building this project! 💪

## 🤝 Contributing

Feel free to contribute! Whether it's Go code improvements, CSS fixes (because I definitely need help there), or new features - all contributions are welcome!

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

## 🙏 Acknowledgments

- [Digi-API](https://digi-api.com/) - For providing the comprehensive Digimon data API
- GitHub Copilot - For being my CSS styling superhero 🦸‍♂️
- The Digimon community - For keeping the digital spirit alive
- The Go community - For creating such an awesome language to learn!

---

*"The courage to take the first step is what separates the dreamer from the doer."* - Tai Kamiya

**Happy coding and digital evolution!** 🌟