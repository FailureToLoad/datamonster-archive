# Datamonster.web

A website for managing Kingdom Death: Monster campaigns.

## Notable frameworks

- [Vite configured for React with Typescript and SWC](https://vitejs.dev/)
- [TailwindCSS](https://tailwindcss.com/)
- [Apollo Client](https://www.apollographql.com/docs/react/)
- [shadcn/ui](https://ui.shadcn.com/docs)
- [Vitest with jsdom](https://vitest.dev/)
- [React Router](https://reactrouter.com/en/main)

## Requirements

A managed [Clerk](https://clerk.com/) instance for authentication.  

## Set Up

Work in progress: I've torn things apart a bit to get the compose working. I need to work backwards towards getting a local workflow that's compatible with the compose workflow.

The key difference is that the web container acts as a reverse proxy for calls to the api container. To mimic this locally I need to put together (and test) instructions for setting up nginx as a reverse proxy.

### Compose

The compose file is set up to use a vault to prove out how this might work in a containzerized deployment. To use it with purely local values, change the environment variables in `api/Dockerfile` to be the following.

```Dockerfile
ENV MODE=dev
ENV KEY="your Clerk secret key"
ENV CONN='host=localhost port=8070 user=app dbname=records password=Password1 sslmode=disable'
ENV CLIENT='dm-web'
```
