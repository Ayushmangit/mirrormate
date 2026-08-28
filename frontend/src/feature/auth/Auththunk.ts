import { createAsyncThunk } from "@reduxjs/toolkit";

import api from "../../api/axiox";
import type {
    LoginResponse,
    LoginUser,
} from "../../types/user";

export const loginUser = createAsyncThunk(
    "auth/login",
    async (user: LoginUser): Promise<LoginResponse> => {
        const response = await api.post("/auth/login", user);

        return response.data;
    }
);