# Web to APK Documentation

## Project Overview
Web to APK is a tool that allows developers to convert their web applications into Android applications. This enables easy distribution and better access to mobile users.

## Features
- Convert web applications to APK format
- Support for both static and dynamic web apps
- Customizable options for app icons and splash screens
- Easy installation and setup

## Installation Instructions
1. **Clone the repository:**
   ```bash
   git clone https://github.com/hanaffy/web-to-apk.git
   cd web-to-apk
   ```
2. **Install dependencies:**
   ```bash
   npm install
   ```
3. **Run the application:**
   ```bash
   npm start
   ```

## API Endpoints
- **POST /convert**: Converts a web URL to APK
  - **Request Body:**
    ```json
    {
      "url": "http://example.com",
      "icon": "path/to/icon.png"
    }
    ```
  - **Response:**
    ```json
    {
      "apk_url": "http://example.com/download/app.apk"
    }
    ```

## Usage Examples
- To convert a web application:
  ```bash
  curl -X POST http://localhost:3000/convert -H 'Content-Type: application/json' -d '{"url": "http://example.com", "icon": "path/to/icon.png"}'
  ```

## Contributing Guidelines
1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them.
4. Push to your branch and create a pull request.

For more information, check the `CONTRIBUTING.md` file in the repository.