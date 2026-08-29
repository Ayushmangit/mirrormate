import axios from "axios";

export const api = axios.create({
	baseURL: import.meta.env.VITE_BASE_URL,

	headers: {
		"Content-Type": "application/json",
	},
});

api.interceptors.request.use((config) => {
	const token = localStorage.getItem("access_token");

	if (token) {
		config.headers.Authorization = `Bearer ${token}`;
	}

	return config;
});

api.interceptors.response.use(
	(response) => response,

	(error) => {
		if (error.response?.status === 401) {
			localStorage.removeItem("access_token");
		}

		return Promise.reject(error);
	},
);
