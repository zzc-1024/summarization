import apiClient from "./apiBase";

export const test = async () => {
  const response = await apiClient.get("/");
  return response.data;
};

export const testUpload = async (formData: FormData) => {
  const response = await apiClient.post("/upload", formData, {
    headers: {
      "Content-Type": "multipart/form-data",
    },
  });
  return response.data;
}