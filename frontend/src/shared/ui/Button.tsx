export const Button = ({  children, onClick, className, type, disabled}: { 
    children: any, 
    onClick?: () => void, 
    className?: string, 
    type?: "submit" | "reset" | "button" | undefined,
    disabled?: boolean,
}) => {
    return <button onClick={onClick} className={className} type={type} disabled={disabled}>
        {children}
    </button>
}