import { FC, useState } from "react";

interface ProductInputProps {
  onSubmit: (
    name: string,
    description: string,
    price: number,
    sku: string
  ) => void;
  onClose: () => void;
}

const ProductInput: FC<ProductInputProps> = ({ onSubmit, onClose }) => {
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [price, setPrice] = useState(0);
  const [sku, setSku] = useState("");
  return (
    <div className="product-details-container">
      <p onClick={onClose} className="x-button">
        X
      </p>
      <input
        type="text"
        maxLength={255}
        placeholder="Name"
        value={name}
        onChange={(e) => setName(e.target.value)}
      />
      <textarea
        maxLength={10000}
        placeholder="Description"
        value={description}
        onChange={(e) => setDescription(e.target.value)}
      />
      <input
        type="number"
        placeholder="Price"
        value={price}
        onChange={(e) => setPrice(+e.target.value)}
      />
      <input
        type="text"
        placeholder="SKU"
        value={sku}
        onChange={(e) => setSku(e.target.value)}
      />
      <button onClick={() => onSubmit(name, description, price, sku)}>
        Add Product
      </button>
    </div>
  );
};
export default ProductInput;
