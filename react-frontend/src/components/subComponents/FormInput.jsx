import React, { useEffect, useState } from "react";

const FormInput = (props) => {
    const [focused, setFocussed] = useState(false)
    const { label, onChange, keyId, inputmeta, wrapperclass, labelclass, error, ...inputProps } = props;

    const handleFocus = () => {
        setFocussed(true);
    };

    return (

        <div className={wrapperclass}>
            <label
                htmlFor={inputProps.id}
                className={labelclass}
            >
                {label}
            </label>

            {inputmeta ? inputmeta.type === "textArea" && (
                <textarea
                    rows={inputmeta.rows} id={inputProps.id}
                    {...inputProps} onChange={(e) => onChange(e, inputProps)} onBlur={handleFocus} focused={focused.toString()}
                ></textarea>
            ) : (
                <input {...inputProps} onChange={(e) => onChange(e, inputProps)} onBlur={handleFocus} focused={focused.toString()} />
            )
            }

            <span className="span-error p-2">{error}</span><div className=" flex flex-wrap" />
        </div>
    );
};

export default FormInput;