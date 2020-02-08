import React from "react";
import MomentUtils from "@date-io/moment";
import {
    MuiPickersUtilsProvider,
    KeyboardDatePicker
} from "@material-ui/pickers";
import { MaterialUiPickersDate } from "@material-ui/pickers/typings/date";

export const DatePickerForm = (props: {
    selectedDate: MaterialUiPickersDate,
    handleDateChange: (date: MaterialUiPickersDate) => void,
}): JSX.Element => {
    return (
        <>
            <MuiPickersUtilsProvider utils={MomentUtils}>
                <KeyboardDatePicker
                    label="dead line"
                    value={props.selectedDate}
                    onChange={props.handleDateChange}
                    format="YYYY/MM/DD"
                />
            </MuiPickersUtilsProvider>
        </>
    );
};