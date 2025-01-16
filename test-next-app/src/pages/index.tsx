export default function Home() {
  return (
    <div>
      <main className=" bg-black flex flex-col gap-8 row-start-2 items-center sm:items-start">
        <h1 dangerouslySetInnerHTML={{ __html: "{{ .Title }}" }} />
        <ul
          className="[&>*]:text-red-500 [&>*]:list-disc"
          dangerouslySetInnerHTML={{
            __html: `{{ range .Books }} <li>{{ . }}</li> {{ end  }}`,
          }}
        />
      </main>
    </div>
  );
}
