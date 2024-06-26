import * as React from "react";
import { fireEvent, render, screen } from "@testing-library/react";

import "@testing-library/jest-dom";

import { ThemeProvider } from "../../Theme";
import DateTimePicker from "../date-time";

afterEach(() => {
  jest.resetAllMocks();
});

const onChange = jest.fn();
test("has padding", () => {
  const { container } = render(
    <ThemeProvider>
      <DateTimePicker value={new Date()} onChange={onChange} />
    </ThemeProvider>
  );

  expect(container.querySelectorAll(".MuiInputBase-adornedEnd")).toHaveLength(1);
  expect(container.querySelector(".MuiInputBase-adornedEnd")).toHaveStyle({
    "padding-right": "14px",
  });
});

test("onChange is called when valid value", () => {
  render(
    <ThemeProvider>
      <DateTimePicker value={new Date()} onChange={onChange} />
    </ThemeProvider>
  );

  expect(screen.getByPlaceholderText("mm/dd/yyyy hh:mm (a|p)m")).toBeVisible();
  fireEvent.change(screen.getByPlaceholderText("mm/dd/yyyy hh:mm (a|p)m"), {
    target: { value: "11/16/2023 02:55 AM" },
  });
  expect(onChange).toHaveBeenCalled();
});

test("onChange is called when valid value and allowEmpty is true", () => {
  render(
    <ThemeProvider>
      <DateTimePicker value={new Date()} onChange={onChange} allowEmpty />
    </ThemeProvider>
  );

  expect(screen.getByPlaceholderText("mm/dd/yyyy hh:mm (a|p)m")).toBeVisible();
  fireEvent.change(screen.getByPlaceholderText("mm/dd/yyyy hh:mm (a|p)m"), {
    target: { value: "11/16/2023 02:55 AM" },
  });
  expect(onChange).toHaveBeenCalled();
});

test("onChange is called when value is empty", () => {
  render(
    <ThemeProvider>
      <DateTimePicker value={new Date()} onChange={onChange} allowEmpty />
    </ThemeProvider>
  );

  expect(screen.getByPlaceholderText("mm/dd/yyyy hh:mm (a|p)m")).toBeVisible();
  fireEvent.change(screen.getByPlaceholderText("mm/dd/yyyy hh:mm (a|p)m"), {
    target: { value: "" },
  });
  expect(onChange).toHaveBeenCalled();
});

test("onChange is not called when invalid value", () => {
  render(
    <ThemeProvider>
      <DateTimePicker value={new Date()} onChange={onChange} />
    </ThemeProvider>
  );

  expect(screen.getByPlaceholderText("mm/dd/yyyy hh:mm (a|p)m")).toBeVisible();
  fireEvent.change(screen.getByPlaceholderText("mm/dd/yyyy hh:mm (a|p)m"), {
    target: { value: "invalid" },
  });
  expect(onChange).not.toHaveBeenCalled();
});

test("onChange is not called when value is empty and empty flag is false", () => {
  render(
    <ThemeProvider>
      <DateTimePicker value={new Date()} onChange={onChange} allowEmpty={false} />
    </ThemeProvider>
  );

  expect(screen.getByPlaceholderText("mm/dd/yyyy hh:mm (a|p)m")).toBeVisible();
  fireEvent.change(screen.getByPlaceholderText("mm/dd/yyyy hh:mm (a|p)m"), {
    target: { value: "" },
  });
  expect(onChange).not.toHaveBeenCalled();
});

test("sets passed value correctly", () => {
  const date = new Date();
  const formattedDMY = new Intl.DateTimeFormat("en-US", {
    month: "2-digit",
    day: "2-digit",
    year: "numeric",
  }).format(date);
  const formattedTime = new Intl.DateTimeFormat("en-US", {
    hour: "2-digit",
    minute: "2-digit",
  }).format(date);
  const formattedDate = `${formattedDMY} ${formattedTime}`;
  render(
    <ThemeProvider>
      <DateTimePicker value={date} onChange={onChange} />
    </ThemeProvider>
  );

  expect(screen.getByPlaceholderText("mm/dd/yyyy hh:mm (a|p)m")).toHaveValue(formattedDate);
});

test("displays label correctly", () => {
  const label = "testing";
  render(
    <ThemeProvider>
      <DateTimePicker value={new Date()} onChange={onChange} label={label} />
    </ThemeProvider>
  );

  expect(screen.getByLabelText(label)).toBeVisible();
});

test("is disabled", () => {
  render(
    <ThemeProvider>
      <DateTimePicker value={new Date()} onChange={onChange} disabled />
    </ThemeProvider>
  );

  expect(screen.getByPlaceholderText("mm/dd/yyyy hh:mm (a|p)m")).toBeDisabled();
});
