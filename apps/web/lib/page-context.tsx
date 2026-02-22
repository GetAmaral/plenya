"use client";

import { createContext, useContext, useState, useCallback, ReactNode } from "react";

export interface PageBreadcrumb {
  label: string;
  href?: string;
}

interface PageHeaderContextValue {
  title: string;
  breadcrumbs: PageBreadcrumb[];
  setPageInfo: (info: { title: string; breadcrumbs?: PageBreadcrumb[] }) => void;
}

const PageHeaderContext = createContext<PageHeaderContextValue>({
  title: "",
  breadcrumbs: [],
  setPageInfo: () => {},
});

export function PageHeaderProvider({ children }: { children: ReactNode }) {
  const [title, setTitle] = useState("");
  const [breadcrumbs, setBreadcrumbs] = useState<PageBreadcrumb[]>([]);

  const setPageInfo = useCallback(
    (info: { title: string; breadcrumbs?: PageBreadcrumb[] }) => {
      setTitle(info.title);
      setBreadcrumbs(info.breadcrumbs ?? []);
    },
    []
  );

  return (
    <PageHeaderContext.Provider value={{ title, breadcrumbs, setPageInfo }}>
      {children}
    </PageHeaderContext.Provider>
  );
}

export function usePageHeader() {
  return useContext(PageHeaderContext);
}
