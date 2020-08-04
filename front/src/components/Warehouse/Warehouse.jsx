import React from 'react';
import Card from 'react-bootstrap/Card';
import { useSelector } from 'react-redux';
import {
  WarehouseEmoji,
  WoodEmoji,
  StoneEmoji,
  ShieldEmoji,
  SwordEmoji,
  MedalEmoji,
} from '../../emojis';

const Warehouse = () => {
  const warehouse = useSelector((state) => state.warehouse)

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
          <WoodEmoji />
          Wood: {warehouse.wood}
        </Card.Text>
        <Card.Text>
        <StoneEmoji />
          Stone: {warehouse.stone}
        </Card.Text>
        <Card.Text>
          <MedalEmoji />
          Gold: {warehouse.gold}
        </Card.Text>
        <Card.Text>
          <SwordEmoji />
          Sword: {warehouse.sword}
        </Card.Text>
        <Card.Text>
          <ShieldEmoji />
          Shield: {warehouse.shield}
        </Card.Text>
      </Card.Body>
    </Card>
  );
}

export default Warehouse;
