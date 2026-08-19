import { configureStore } from '@reduxjs/toolkit'
import authReducer from '../feature/auth/authslice'

export const store = configureStore({
  reducer: {
    auth: authReducer,
  },
})
