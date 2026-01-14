interface Props extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  children: React.ReactNode;
  className?: string;
}

function Button({ children, className, ...props }: Props) {
  // Not quite good in not using 'cn' but I'm too lazy to import
  return (
    <button
      className={`px-4 py-2 bg-slate-400 text-white rounded-lg font-semibold hover:bg-slate-600/90 duration-300 ${className}`}
      {...props}
    >
      {children}
    </button>
  );
}

export { Button };
