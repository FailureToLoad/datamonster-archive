# Datamonster.web

A website for managing Kingdom Death: Monster campaigns.

## Notable frameworks

- [Vite configured for React with Typescript and SWC](https://vitejs.dev/)
- [TailwindCSS](https://tailwindcss.com/)
- [Tanstack Query](https://tanstack.com/query/latest/docs/framework/react/overview)
- [shadcn/ui](https://ui.shadcn.com/docs)
- [Vitest with jsdom](https://vitest.dev/)
- [React Router](https://reactrouter.com/en/main)

## Related Projects

 [datamonster.api](https://github.com/FailureToLoad/datamonster.api) - The api this site is designed to work against.  
 [datamonster.records](https://github.com/FailureToLoad/datamonster.records) - A containerized postgres database the api uses for persistence.  

## Requirements

A managed [Clerk](https://clerk.com/) instance for authentication.  

## Set Up

- Create your own .env.local file by using the env.local.example file as a base.
- Clone, set up, build, and run [datamonster.records](https://github.com/FailureToLoad/datamonster.records).
- Clone, set up, build and run [datamonster.api](https://github.com/FailureToLoad/datamonster.api).
