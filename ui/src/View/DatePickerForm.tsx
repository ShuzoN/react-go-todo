import React, { Fragment, useState } from "react";
import MomentUtils from "@date-io/moment";
import moment from 'moment';
import {
    MuiPickersUtilsProvider,
    KeyboardDatePicker
} from "@material-ui/pickers";
import { MaterialUiPickersDate } from "@material-ui/pickers/typings/date";




export const DatePickerForm = (): JSX.Element => {
    const [selectedDate, setSelectedDate] = useState<MaterialUiPickersDate>(moment(new Date()));

    console.log(selectedDate);

    const handleDateChange = (date: MaterialUiPickersDate) => {
        setSelectedDate(date);
    };

    return (
        <Fragment>
            <MuiPickersUtilsProvider utils={MomentUtils}>
                <KeyboardDatePicker
                    label="dead line"
                    value={selectedDate}
                    onChange={handleDateChange}
                    format="YYYY/MM/DD"
                />
            </MuiPickersUtilsProvider>
        </Fragment>
    );
};