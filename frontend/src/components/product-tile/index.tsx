import { FC } from "react";

interface ProductTileProps {
  name: string;
  price: number;
  onClick: () => void;
  onDelete: () => void;
}

const ProductTile: FC<ProductTileProps> = ({
  name,
  price,
  onClick,
  onDelete,
}) => {
  return (
    <div onClick={onClick} className="product-tile-container">
      <h1>{name}</h1>
      <div className="row">
        <p className="price">Price: {price}</p>
        <p onClick={onDelete} className="x-button">
          X
        </p>
      </div>
    </div>
  );
};
export default ProductTile;
