import { addProduct, getProducts } from "./functions";
import { useEffect, useState } from "react";

import { Product } from "./models/product";
import ProductInput from "./components/product-input";
import ProductList from "./components/product-list";

function App() {
  const [showProductInput, setShowProductInput] = useState(false);
  const [products, setProducts] = useState<Product[]>([]);
  useEffect(() => {
    (async () => {
      const products = await getProducts();
      setProducts(products);
    })();
  }, []);

  return (
    <div className="flex-column">
      <div className="row">
        <button
          className="add-product-button"
          onClick={() => setShowProductInput(true)}
        >
          Add Product
        </button>
        <button
          className="refresh-button"
          onClick={async () => {
            const products = await getProducts();
            setProducts(products);
          }}
        >
          Refresh
        </button>
      </div>

      {showProductInput && (
        <ProductInput
          onSubmit={async (name, description, price, sku) => {
            const newProduct = {
              name,
              description,
              price,
              sku,
            };
            await addProduct(newProduct);
            const products = await getProducts();
            setProducts(products);
            setShowProductInput(false);
          }}
          onClose={() => setShowProductInput(false)}
        />
      )}

      <ProductList products={products} setProducts={setProducts} />
    </div>
  );
}

export default App;
