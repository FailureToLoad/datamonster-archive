import {Link, useRouteError, isRouteErrorResponse} from 'react-router-dom';

const ErrorPage = () => {
  const error = useRouteError();
  if (isRouteErrorResponse(error)) {
    return (
      <div className="mt-28 flex h-screen w-full flex-col items-center">
        <h1 className="text-neutral-grayish-blue text-8xl">
          Error{' '}
          <span className="text-primary-soft-red text-8xl">{error.status}</span>
        </h1>
        <p className="text-neutral-grayish-blue my-10 text-4xl">
          {error.statusText}
        </p>
        <Link
          to={'/welcome'}
          className="bg-neutral-dark-blue text-neutral-white rounded-md p-2"
        >
          Back Home
        </Link>
      </div>
    );
  } else {
    return (
      <div className="mt-28 flex h-screen w-full flex-col items-center">
        <h1 className="text-neutral-grayish-blue text-8xl">Error </h1>
        <p className="text-neutral-grayish-blue my-10 text-4xl">
          <i>
            {(error as Error)?.message ||
              (error as {statusText?: string})?.statusText}
          </i>
        </p>
        <div className="size-1/2 w-1/2 text-wrap">
          <code>{(error as Error).stack}</code>
        </div>
        <Link
          to={'/welcome'}
          className="bg-neutral-dark-blue text-neutral-white rounded-md p-2"
        >
          Back Home
        </Link>
      </div>
    );
  }
};

export default ErrorPage;
