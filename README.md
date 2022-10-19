# Salmon

Managing libvirt virtual machines is a pain. With salmon you can manage
your libvirt virtual machines with a stable and centralized rest api
making it awesome for scripting and automation.

# Node

A node is basically just the virtualization server that is virtualizing
all the virtual machines. A machine can also be referred to as a
subnode.

# Features

1. [?] A agent which runs on the virtualization server to collect
   data about the host (node) and its machines.
2. [?] A api which can be used for inspecting the machines and
   creating new ones. 
3. [?] A web interface for visual representation of your servers mapped
   by salmon.
4. [] A page in the web interface that visualizes the network from your
   virtual machines to the router for those of us who are forced in to
   using shit routers and shit router software.
5. [] Integration against DNS servers to automatically map both nodes
   and virtual machines to a FQDN. This automatically replaces the need
   for mDNS in the homelab or datacenter. Structured by the virtual
   machines hostname and the nodes hostname and a chosen subdomain.
   Example: *cyan.green.lab.eskpil.com*

