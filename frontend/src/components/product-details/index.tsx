import { FC } from "react";

interface ProductDetailsProps {
  name: string;
  description: string;
  price: number;
  sku: string;
  onClose: () => void;
}

const ProductDetails: FC<ProductDetailsProps> = ({
  name,
  description,
  price,
  sku,
  onClose,
}) => {
  return (
    <div className="product-details-container">
      <div className="row">
        <h1>{name}</h1>
        <p onClick={onClose} className="x-button">
          X
        </p>
      </div>
      <p>{description}</p>
      <p className="price">Price: {price}</p>
      <p className="sku">SKU: {sku}</p>
    </div>
  );
};
export default ProductDetails;
