export const deleteProduct = async (id: number) => {
  try {
    const res = await fetch(window.global.api_location + "/products/" + id, {
      method: "DELETE",
      mode: "cors",
    });
    console.log(res);
  } catch (error) {
    console.error(error);
  }
};
