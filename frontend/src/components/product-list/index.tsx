import { FC, useState } from "react";
import { deleteProduct, getProducts } from "../../functions";

import { Product } from "../../models/product";
import ProductDetails from "../product-details";
import ProductTile from "../product-tile";

interface ProductListProps {
  products: Product[];
  setProducts: (products: Product[]) => void;
}

const ProductList: FC<ProductListProps> = ({ products, setProducts }) => {
  const [selectedProduct, setSelectedProduct] = useState<Product | null>(null);

  return (
    <div className="flex-column">
      {products.map((product) => (
        <ProductTile
          key={product.id}
          name={product.name}
          price={product.price}
          onClick={() => setSelectedProduct(product)}
          onDelete={async () => {
            await deleteProduct(product.id);
            const products = await getProducts();
            setSelectedProduct(null);
            setProducts(products);
          }}
        />
      ))}
      {selectedProduct && (
        <ProductDetails
          name={selectedProduct.name}
          description={selectedProduct.description}
          price={selectedProduct.price}
          sku={selectedProduct.sku}
          onClose={() => setSelectedProduct(null)}
        />
      )}
    </div>
  );
};
export default ProductList;
