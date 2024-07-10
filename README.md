## Simple Looker Embed SDK Demo (Go)

This demo showcases basic Go code for using Signed Embed to render
a Looker Dashboard locally.

It utilizes Docker containers for a streamlined setup.

**Prerequisites:**

* Docker and Docker Compose installed ([https://docs.docker.com/](https://docs.docker.com/))

**Mandatory Environment Variable:**

* `LOOKER_SERVER_URL`: The URL of your Looker instance.

**Steps:**

1. **Generate Looker API Credentials:**
   - Access your Looker admin panel and navigate to your user settings.
   - Click "API Keys" and generate a new Client ID and Client Secret. Copy them for later use.

2. **Set Up Environment Variables:**

   **Option 1: Using the Provided .env.sample File**
     - Copy the provided `.env.sample` file to `.env`.
     - Edit the `.env` file and replace placeholders with your Looker credentials:
        - `LOOKER_CLIENT_ID`
        - `LOOKER_CLIENT_SECRET`
        - `LOOKER_SERVER_URL`
 
3. **Run the Demo:**
   - Open a terminal and navigate to the project directory.
   - Run `docker-compose up --build` to start the demo environment.
   - Access the demo application in your browser at `http://localhost:8080/`.

