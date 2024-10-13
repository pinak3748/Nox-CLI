package content

import (
	"strings"
)

func GenerateSliceContent(pageName string) string {
	template := `import { PayloadAction, createSlice } from '@reduxjs/toolkit';
import { {pageName}Item } from '../../../constants/types';
import {
  create{pageName},
  delete{pageName},
  get{pageName},
  get{pageName}s,
  update{pageName},
} from './{pageName}Action';

interface {pageName}State {
  items: {pageName}Item[];
  currentItem: {pageName}Item | null;
  loading: boolean;
  error: string | null;
  selectedItem: {pageName}Item | null;
}

const initialState: {pageName}State = {
  items: [],
  currentItem: null,
  loading: false,
  error: null,
  selectedItem: null,
};

const {pageName}Slice = createSlice({
  name: '{pageName}',
  initialState,
  reducers: {
    setSelected{pageName}: (state, action: PayloadAction<{pageName}Item | null>) => {
      state.selectedItem = action.payload;
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(create{pageName}.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(create{pageName}.fulfilled, (state, action: PayloadAction<{pageName}Item>) => {
        state.loading = false;
        state.items.push(action.payload);
      })
      .addCase(create{pageName}.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload as string;
      })
      .addCase(get{pageName}s.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(get{pageName}s.fulfilled, (state, action: PayloadAction<{pageName}Item[]>) => {
        state.loading = false;
        state.items = action.payload;
      })
      .addCase(get{pageName}s.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload as string;
      })
      .addCase(get{pageName}.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(get{pageName}.fulfilled, (state, action: PayloadAction<{pageName}Item>) => {
        state.loading = false;
        state.currentItem = action.payload;
      })
      .addCase(get{pageName}.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload as string;
      })
      .addCase(update{pageName}.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(update{pageName}.fulfilled, (state, action: PayloadAction<{pageName}Item>) => {
        state.loading = false;
        const index = state.items.findIndex(
          (item) => item.id === action.payload.id
        );
        if (index !== -1) {
          state.items[index] = action.payload;
        }
        if (state.currentItem && state.currentItem.id === action.payload.id) {
          state.currentItem = action.payload;
        }
      })
      .addCase(update{pageName}.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload as string;
      })
      .addCase(delete{pageName}.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(delete{pageName}.fulfilled, (state, action: PayloadAction<string>) => {
        state.loading = false;
        state.items = state.items.filter(
          (item) => item.id !== action.payload
        );
        if (state.currentItem && state.currentItem.id === action.payload) {
          state.currentItem = null;
        }
      })
      .addCase(delete{pageName}.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload as string;
      });
  },
});

export const { setSelected{pageName} } = {pageName}Slice.actions;
export default {pageName}Slice.reducer;
`

	return strings.ReplaceAll(template, "{pageName}", pageName)
}
