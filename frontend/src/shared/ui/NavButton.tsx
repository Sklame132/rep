
export const NavButton = ({children, className, isActive, onClick}: {children: any, className: string, isActive?: boolean, onClick: () => void}) => {
    return <>
        <li onClick={onClick} className={[
            [className, "__button"].join(''), 
            isActive ? [className, "__button-active"].join('') : '']
            .join(' ')}>
            <a className={[className, "__link"].join('')}>
                {children}
            </a>
        </li>
    </>
}