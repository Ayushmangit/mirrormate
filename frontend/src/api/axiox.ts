import axiox from "axios";

const api = axiox.create({
  baseURL: "http://localhost:8080/api",
});
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem("token");
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  }
);
export default api;