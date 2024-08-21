import {Play} from '@phosphor-icons/react';
import {Link} from 'react-router-dom';
import {SettlementId} from '@/lib/types/settlements';
import {Card, CardContent, CardHeader, CardTitle} from '@/components/ui/card';
import {Button} from '@/components/ui/button';

export function SettlementCard({settlement}: {settlement: SettlementId}) {
  const link = '/settlements/' + settlement.id;
  return (
    <Card>
      <CardHeader>
        <CardTitle>{settlement.name}</CardTitle>
      </CardHeader>
      <CardContent>
        <div className="flex flex-row justify-between">
          <div>
            <Link to={link}>
              <Button variant="ghost" size="icon">
                <Play className="h-6 w-6" />
              </Button>
            </Link>
          </div>
        </div>
      </CardContent>
    </Card>
  );
}
