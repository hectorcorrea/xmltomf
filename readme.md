# xmltomf

Command line utility to convert a TigerData XML document to a Mediaflux document


**this is a work in progress**

# Sample of usage

Run the following command to process the `resourceTiny.xml` file and generate a Mediaflux document representation:

```
./xmltomf -file resourceTiny.xml
```

Outputs something along the lines of:

```
:meta < \
  tigerdataX:resourceDoc < \
    :resource -resourceClass "Project" -resourceID "10.34770/az09-0001" -resourceIDTYpe "DOI" < \
      :projectID -projectIDType "DOI" "10.34770/az09-0001" \
    >
  >
>
```

# Download
Download the executable for your operating system from: