import { createSlice } from "@reduxjs/toolkit";

import { loginUser, logoutUser, registerUser } from "./authThunk";
import type { User } from "../../types/user";

interface AuthState {
	userData: User | null;
	loading: boolean;
	accessToken: string | null;
	error: string | null;
}

const accessToken = localStorage.getItem("access_token");

const initialState: AuthState = {
	userData: null,
	loading: false,
	accessToken,
	error: null,
};

const authSlice = createSlice({
	name: "auth",

	initialState,

	reducers: {
		clearAuthError: (state) => {
			state.error = null;
		},
	},

	extraReducers: (builder) => {
		builder

			// REGISTER

			.addCase(registerUser.pending, (state) => {
				state.loading = true;
				state.error = null;
			})

			.addCase(registerUser.fulfilled, (state, action) => {
				state.loading = false;
				state.error = null;
				state.userData = action.payload.user;
			})

			.addCase(registerUser.rejected, (state, action) => {
				state.loading = false;

				state.error = action.payload ?? "Registration failed";
			})

			// LOGIN

			.addCase(loginUser.pending, (state) => {
				state.loading = true;
				state.error = null;
			})

			.addCase(loginUser.fulfilled, (state, action) => {
				state.loading = false;
				state.error = null;

				state.userData = action.payload.user;

				state.accessToken = action.payload.access_token;

				localStorage.setItem("access_token", action.payload.access_token);
			})

			.addCase(loginUser.rejected, (state, action) => {
				state.loading = false;

				state.error = action.payload ?? "Login failed";
			})

			// LOGOUT

			.addCase(logoutUser.fulfilled, (state) => {
				state.userData = null;
				state.accessToken = null;
				state.loading = false;
				state.error = null;
			});
	},
});

export const { clearAuthError } = authSlice.actions;

export default authSlice.reducer;
