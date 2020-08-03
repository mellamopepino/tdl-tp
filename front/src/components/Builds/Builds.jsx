import React from 'react';
import Card from 'react-bootstrap/Card';
import { useSelector } from 'react-redux';
import {
  SwordEmoji,
  WorkerEmoji,
  ShieldEmoji,
  WrenchEmoji,
} from '../../emojis/'

const Builds = () => {
  const workers = useSelector((state) => state.workers)

  return (
    <Card className="builds">
      <Card.Header>
        <h2>
          <WrenchEmoji />
          Builds
        </h2>
      </Card.Header>
      <Card.Body>
        <Card.Text>
          <SwordEmoji/> <WorkerEmoji/> [{workers.sword} working...]
        </Card.Text>
        <Card.Text>
          <ShieldEmoji/> <WorkerEmoji/> [{workers.shield} working...]
        </Card.Text>
      </Card.Body>
    </Card>
  );
}

export default Builds;
