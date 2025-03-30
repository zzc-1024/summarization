import apiClient from "./apiBase";

export const test = async () => {
  const response = await apiClient.get("/");
  return response.data;
};

export const getTasks = async () => {
  const response = await apiClient.get("/task");
  return response.data;
}

export const getTask = async (id: string) => {
  const response = await apiClient.get(`/task/${id}`);
  return response.data;
}

export const uploadFile = async (formData: FormData) => {
  const response = await apiClient.post("/task/file", formData, {
    headers: {
      "Content-Type": "multipart/form-data",
    },
  });
  return response.data;
}

export const abstractTask = async (taskId: string) => {
  const response = await apiClient.post(`/task/${taskId}/abstract`);
  return response.data;
}
