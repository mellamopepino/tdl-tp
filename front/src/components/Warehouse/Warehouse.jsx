import React from 'react';
import Card from 'react-bootstrap/Card';
import {
  WarehouseEmoji,
  WoodEmoji,
  StoneEmoji,
} from '../../emojis';

const Warehouse = () => {
  return (
    <Card className="warehouse">
      <Card.Header>
        <h2>
          <WarehouseEmoji />
          Warehouse
        </h2>
      </Card.Header>
      <Card.Body>
        <Card.Text>
          <p>
            <WoodEmoji />
            Wood: 10
          </p>
          <p>
            <StoneEmoji />
            Stone: 20
          </p>
        </Card.Text>
      </Card.Body>
    </Card>
  );
}

export default Warehouse;
