import React from "react";

const AccountSelector = ({
  accounts,
  handleChange,
  value,
  message,
  reference,
}) => (
  <select name="" id="" onChange={handleChange} value={value} ref={reference}>
    {accounts.length &&
      accounts.map((acc, idx) => {
        return idx === 0 ? (
          <>
            <option value={message} key={"no-value"}>
              {message}
            </option>
            <option value={acc} key={idx + 1}>
              {acc}
            </option>
          </>
        ) : (
          <option value={acc} key={idx + 1}>
            {acc}
          </option>
        );
      })}
  </select>
);

export default AccountSelector;
