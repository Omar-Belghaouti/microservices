import { Product } from "../models/product";

export const getProducts = async () => {
  try {
    const res = await fetch(window.global.api_location + "/products/");
    const json = await res.json();
    return json;
  } catch (error) {
    console.error(error);
  }
};
