import { Product } from "../models/product";

export const addProduct = async (product: Partial<Product>) => {
  try {
    const res = await fetch(window.global.api_location + "/products/", {
      method: "POST",
      body: JSON.stringify(product),
    });
    const json = await res.json();
    console.log(json);
  } catch (error) {
    console.error(error);
  }
};
