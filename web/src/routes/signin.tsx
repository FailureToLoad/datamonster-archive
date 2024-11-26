// src/routes/signin.tsx
import {useEffect} from 'react';
import {useNavigate} from 'react-router-dom';
import {useAuth} from '@/auth/hooks';
import {Button} from '@/components/ui/button';
import {Card, CardContent, CardHeader, CardTitle} from '@/components/ui/card';

export default function SignInPage() {
  const {signIn, isAuthenticated, isLoading} = useAuth();
  const navigate = useNavigate();

  useEffect(() => {
    if (isAuthenticated) {
      navigate('/settlements', {replace: true});
    }
  }, [isAuthenticated, navigate]);

  if (isLoading) {
    return <div>Loading...</div>;
  }

  return (
    <div className="flex min-h-screen items-center justify-center">
      <Card className="w-[350px]">
        <CardHeader>
          <CardTitle>Sign In to Datamonster</CardTitle>
        </CardHeader>
        <CardContent>
          <Button onClick={signIn} className="w-full" variant="outline">
            <img src="/google.svg" alt="Google" className="mr-2 h-4 w-4" />
            Sign in with Google
          </Button>
        </CardContent>
      </Card>
    </div>
  );
}
