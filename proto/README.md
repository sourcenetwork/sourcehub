# Proto

Protobuff definitions for SourceHub.

# Buf
SourceHub makes use of buf to build and manage its proto types.

# Acp Module
The ACP module depends on `zanzi`'s proto definitions.
The dependencies are linked through the `buf.work.yaml` file in the project root.
The workspace file references the buf module inside the source-zanzibar repository.
The source-zanzibar repository is fetched as a git-submodule, the submodule was chosen because source-zanzibar proto types are still evolving, therefore it doesn't make sense to publish them in the buf repository.
