# Simple Looker Embed SDK Demo (Go)

This demo showcases basic Go code for using Signed Embed to render
a Looker Dashboard locally.

It utilizes Docker containers for a streamlined setup.

## Prerequisites:

* Docker and Docker Compose installed ([https://docs.docker.com/](https://docs.docker.com/))

## Steps:

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


## Deploy to Cloud Run

This option will allow you to run the project as a hosted Cloud Run app.
This requires that all previous steps were completed, including the
`.env` file configuration part with the proper credentials.

1. **Source the environment variables**

```bash
source .env
```

2. **Create secres on Secrets Manager**

```bash
printf "${LOOKERSDK_CLIENT_ID}" | gcloud secrets create LOOKERSDK_CLIENT_ID --data-file=-
printf "${LOOKERSDK_CLIENT_SECRET}" | gcloud secrets create LOOKERSDK_CLIENT_SECRET --data-file=-
```

3. **Deploy the service to Cloud Run**

```bash
gcloud run deploy looker-embed-demo \
   --set-env-vars="LOOKERSDK_BASE_URL=$LOOKERSDK_BASE_URL" \
   --set-secrets="LOOKERSDK_CLIENT_ID=LOOKERSDK_CLIENT_ID:latest,LOOKERSDK_CLIENT_SECRET=LOOKERSDK_CLIENT_SECRET:latest" \
   --region us-central1 \
   --source .
```