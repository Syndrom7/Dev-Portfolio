# Dev Portfolio

Welcome to the repository for my personal development portfolio, a showcase of my skills, projects, and experiences as a web developer.

## Overview

This portfolio is built with Go and front-end technologies (HTML, CSS, JavaScript, Alpine.js, and Tailwind CSS). It is designed to be minimalistic, responsive, and easily navigable, presenting a clean and professional interface.

## Features

- **Dynamic Content**: Server-side rendering with Go templates for dynamic content.
- **Responsive Design**: Tailwind CSS is used to ensure the portfolio looks great on all devices.
- **Interactive UI**: Alpine.js provides the reactivity and interactivity of the user interface.
- **Portfolio Projects**: A showcase of my latest work and projects.
- **Contact Form**: Allows visitors to send messages or inquiries directly through the website.

## Local Development

To run this project locally, follow these steps:

1. Clone the repository:
   ```bash
   git clone https://github.com/Syndrom7/dev-portfolio.git
    ```

2. Navigate to the project directory:
   ```bash
   cd dev-portfolio
   ```

3. Run the Go server:
   ```bash
   go run main.go
   ```

This will start the server on `localhost:8081`

## Docker Setup

To run this project using docker, follow these steps:

1. Build the Docker image
   ```bash
   docker build -t dev_portfolio .
   ```

2. Run the Docker container
   ```bash
   docker run -p 8081:8081 dev_portfolio
   ```

3. Open your browser and visit `http://localhost:8081`.

## Structure
- `src/`: Contains all the source files, including HTML templates, CSS, and JavaScript.
- `src/partials/`: Reusable HTML components like headers and footers.
- `assets/`: Static files like images and fonts.
- `main.go`: The Go server entry point.