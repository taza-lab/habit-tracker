"use client"

import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";
import { ThemeProvider, CssBaseline } from '@mui/material';
import { usePathname } from "next/navigation";
import theme from './theme';
import Layout from "@/components/Layout";
import { PointProvider } from '@/context/PointContext';
import { AlertProvider } from '@/context/AlertContext';

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {

  const pathname = usePathname();
  const noLayoutPaths = ["/login", "/signup"];
  const isLoginedPath = !noLayoutPaths.includes(pathname);

  return (
    <html lang="">
      <head>
        <link
          href="https://cdn.jsdelivr.net/npm/bootstrap-icons@1.10.5/font/bootstrap-icons.css"
          rel="stylesheet"
        />
      </head>
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased`}
      >
        <ThemeProvider theme={theme}>
          <CssBaseline />
          <PointProvider>
            <AlertProvider>
              <Layout showLoginedLayout={isLoginedPath}>{children}</Layout>
            </AlertProvider>
          </PointProvider>
        </ThemeProvider>
      </body>
    </html>
  );
}
