import { createAsyncThunk } from "@reduxjs/toolkit";
import axios from "axios";

import { api } from "../../api/axios";
import type {
	LoginPayload,
	LoginResponse,
	RegisterPayload,
	RegisterResponse,
} from "../../types/auth";

function getErrorMessage(error: unknown): string {
	if (axios.isAxiosError(error)) {
		return (
			error.response?.data?.error ?? "Something went wrong. Please try again."
		);
	}

	return "Something went wrong. Please try again.";
}

export const registerUser = createAsyncThunk<
	RegisterResponse,
	RegisterPayload,
	{ rejectValue: string }
>("auth/register", async (data, { rejectWithValue }) => {
	try {
		const response = await api.post("/auth/register", data);

		const { user } = response.data.data;

		return {
			user,
		};
	} catch (error) {
		return rejectWithValue(getErrorMessage(error));
	}
});

export const loginUser = createAsyncThunk<
	LoginResponse,
	LoginPayload,
	{ rejectValue: string }
>("auth/login", async (data, { rejectWithValue }) => {
	try {
		const response = await api.post("/auth/login", data);

		const { user, access_token } = response.data.data;

		return {
			user,
			access_token,
		};
	} catch (error) {
		return rejectWithValue(getErrorMessage(error));
	}
});

export const logoutUser = createAsyncThunk(
	"auth/logout",
	async (_, { rejectWithValue }) => {
		try {
			localStorage.removeItem("access_token");

			return null;
		} catch {
			return rejectWithValue("Logout failed");
		}
	},
);
